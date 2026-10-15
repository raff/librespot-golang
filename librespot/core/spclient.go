package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/librespot-org/librespot-golang/Spotify"
	"google.golang.org/protobuf/encoding/protowire"
)

const (
	extendedMetadataUrl = "https://spclient.wg.spotify.com/extended-metadata/v0/extended-metadata"

	// ExtensionKind AUDIO_FILES
	extensionAudioFiles = 5
)

// TrackAudioFiles returns the audio files (file id and format, including lossless FLAC ones) available for a
// track, using Spotify's extended-metadata service. The legacy hm://metadata/4 endpoint doesn't list lossless
// files. trackUri is a "spotify:track:<base62 id>" URI. It requires credentials usable with login5 (see
// Session.SpclientToken).
func (s *Session) TrackAudioFiles(trackUri string) ([]*Spotify.AudioFile, error) {
	token, err := s.SpclientToken()
	if err != nil {
		return nil, err
	}

	// BatchedEntityRequest { header(1) { catalogue(2) }, entity_request(2) { entity_uri(1), query(2) { extension_kind(1) } } }
	str := func(num protowire.Number, v string) []byte {
		return protowire.AppendString(protowire.AppendTag(nil, num, protowire.BytesType), v)
	}
	msg := func(num protowire.Number, v []byte) []byte {
		return protowire.AppendBytes(protowire.AppendTag(nil, num, protowire.BytesType), v)
	}

	query := protowire.AppendVarint(protowire.AppendTag(nil, 1, protowire.VarintType), extensionAudioFiles)
	body := msg(1, str(2, "premium"))
	body = append(body, msg(2, append(str(1, trackUri), msg(2, query)...))...)

	req, err := http.NewRequest("POST", extendedMetadataUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/protobuf")
	req.Header.Set("Accept", "application/protobuf")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("extended-metadata failed: status %d", resp.StatusCode)
	}

	return parseAudioFiles(data)
}

// eachField iterates over the fields of a protobuf message
func eachField(b []byte, fn func(num protowire.Number, typ protowire.Type, val []byte, v uint64) error) error {
	for len(b) > 0 {
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return protowire.ParseError(n)
		}
		b = b[n:]

		switch typ {
		case protowire.BytesType:
			v, n := protowire.ConsumeBytes(b)
			if n < 0 {
				return protowire.ParseError(n)
			}
			if err := fn(num, typ, v, 0); err != nil {
				return err
			}
			b = b[n:]
		case protowire.VarintType:
			v, n := protowire.ConsumeVarint(b)
			if n < 0 {
				return protowire.ParseError(n)
			}
			if err := fn(num, typ, nil, v); err != nil {
				return err
			}
			b = b[n:]
		default:
			n := protowire.ConsumeFieldValue(num, typ, b)
			if n < 0 {
				return protowire.ParseError(n)
			}
			b = b[n:]
		}
	}
	return nil
}

// parseAudioFiles extracts the files from a BatchedExtensionResponse:
//
//	{ extended_metadata(2) { extension_data(3) { data(3) { type_url(1), value(2) = AudioFilesExtensionResponse } } } }
func parseAudioFiles(data []byte) ([]*Spotify.AudioFile, error) {
	var files []*Spotify.AudioFile

	// descend follows a path of length-delimited field numbers, calling fn on the innermost message
	var descend func(b []byte, path []protowire.Number, fn func([]byte) error) error
	descend = func(b []byte, path []protowire.Number, fn func([]byte) error) error {
		if len(path) == 0 {
			return fn(b)
		}
		return eachField(b, func(num protowire.Number, typ protowire.Type, val []byte, _ uint64) error {
			if num == path[0] && typ == protowire.BytesType {
				return descend(val, path[1:], fn)
			}
			return nil
		})
	}

	err := descend(data, []protowire.Number{2, 3, 3, 2}, func(b []byte) error {
		return parseAudioFilesResponse(b, &files)
	})
	return files, err
}

// AudioFilesExtensionResponse { files(1) { file(1) { file_id(1), format(2) }, ... } }
func parseAudioFilesResponse(b []byte, files *[]*Spotify.AudioFile) error {
	return eachField(b, func(num protowire.Number, typ protowire.Type, val []byte, v uint64) error {
		if num != 1 || typ != protowire.BytesType {
			return nil
		}
		return eachField(val, func(n2 protowire.Number, t2 protowire.Type, v2 []byte, _ uint64) error {
			if n2 != 1 || t2 != protowire.BytesType {
				return nil
			}

			f := &Spotify.AudioFile{}
			err := eachField(v2, func(n3 protowire.Number, t3 protowire.Type, v3 []byte, v uint64) error {
				switch {
				case n3 == 1 && t3 == protowire.BytesType:
					f.FileId = append([]byte(nil), v3...)
				case n3 == 2 && t3 == protowire.VarintType:
					f.Format = Spotify.AudioFile_Format(v).Enum()
				}
				return nil
			})
			if err != nil {
				return err
			}
			*files = append(*files, f)
			return nil
		})
	})
}
