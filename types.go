package main

import mapset "github.com/deckarep/golang-set/v2"

type Mod struct {
	Id string
	// Versions map[string]ModVersion
	DepsMap map[string]mapset.Set[string]
}

func NewMod() Mod {
	var mod Mod
	// mod.Versions = make(map[string]ModVersion)
	mod.DepsMap = make(map[string]mapset.Set[string])
	return mod
}

type ModVersion struct {
	Version string
	DepsMap map[string]mapset.Set[string]
}

func NewVersion() ModVersion {
	var version ModVersion
	version.DepsMap = make(map[string]mapset.Set[string])
	return version
}
