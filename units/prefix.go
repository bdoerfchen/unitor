package units

import (
	"math"

	"github.com/bdoerfchen/unitor"
)

var PrefixesSI = unitor.NewUnit("SI Prefixes").
	WithPrefix("Q", 1e30).
	WithPrefix("R", 1e27).
	WithPrefix("Y", 1e24).
	WithPrefix("Z", 1e21).
	WithPrefix("E", 1e18).
	WithPrefix("P", 1e15).
	WithPrefix("T", 1e12).
	WithPrefix("G", 1e9).
	WithPrefix("M", 1e6).
	WithPrefix("k", 1e3).
	WithPrefix("h", 1e2).
	WithPrefix("da", 1e1).
	WithPrefix("d", 1e-1).
	WithPrefix("c", 1e-2).
	WithPrefix("m", 1e-3).
	WithPrefix("µ", 1e-6, "u").
	WithPrefix("n", 1e-9).
	WithPrefix("p", 1e-12).
	WithPrefix("f", 1e-15).
	WithPrefix("a", 1e-18).
	WithPrefix("z", 1e-21).
	WithPrefix("y", 1e-24).
	WithPrefix("r", 1e-27).
	WithPrefix("q", 1e-30).
	Unit()

var PrefixesBinary = unitor.NewUnit("Binary Prefixes").
	WithPrefix("Ki", math.Pow(1024, 0)).
	WithPrefix("Mi", math.Pow(1024, 2)).
	WithPrefix("Gi", math.Pow(1024, 3)).
	WithPrefix("Ti", math.Pow(1024, 4)).
	WithPrefix("Pi", math.Pow(1024, 5)).
	WithPrefix("Ei", math.Pow(1024, 6)).
	WithPrefix("Zi", math.Pow(1024, 7)).
	WithPrefix("Yi", math.Pow(1024, 8)).
	WithPrefix("Ri", math.Pow(1024, 9)).
	WithPrefix("Qi", math.Pow(1024, 10)).
	Unit()
