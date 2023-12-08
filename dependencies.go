package main

import (
	"archive/zip"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

var ()

func ReadMods(dir string) (map[string]Mod, error) {

	var mods map[string]Mod = make(map[string]Mod)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jar") {
			if verbose {
				log.Printf("file.Name(): %v\n", file.Name())
			}
			var modId string
			var mod Mod

			// Open jar for reading
			r, err := zip.OpenReader(dir + "/" + file.Name())
			if err != nil {
				continue
			}
			defer r.Close()

			// Open json file for reading
			rc, err := r.Open("fabric.mod.json")
			if err != nil {
				continue
			}

			// Read json data
			byteValue, _ := io.ReadAll(rc)
			data := make(map[string]interface{})
			json.Unmarshal(byteValue, &data)

			if id, ok := data["id"].(string); ok {
				modId = id
			} else {
				continue
			}

			if _, ok := mods[modId]; ok {
				mod = mods[modId]
			} else {
				mod = NewMod()
				mod.Id = modId
			}

			if deps, ok := data["depends"].(map[string]interface{}); ok {
				for k, v := range deps {
					if _, ok := mods[k]; !ok {
						var depMod Mod = NewMod()
						depMod.Id = k
						mods[k] = depMod
						if verbose {
							log.Printf("depmod: %v\n", k)
						}
					}
					// TODO: there's probably a better way of doing this
					if _, ok := mod.DepsMap[k]; !ok {
						mod.DepsMap[k] = mapset.NewSet[string]()
					}

					switch rv := v.(type) {
					case string:
						mod.DepsMap[k].Add(rv)
					case []interface{}:
						for _, lv := range rv {
							mod.DepsMap[k].Add(lv.(string))
						}
					}
				}
			}
			mods[mod.Id] = mod
		}
	}

	return mods, nil
}
