package unitconverter

import "errors"

func Convert(value float64, from string, to string) (float64, error) {
    switch from + "_" + to {
    case "km_m":
        return value * 1000, nil
    case "m_km":
        return value / 1000, nil
    case "kg_g":
        return value * 1000, nil
    case "g_kg":
        return value / 1000, nil
    case "c_f":
        return (value*9/5 + 32), nil
    case "f_c":
        return (value-32)*5/9, nil
    default:
        return 0, errors.New("unsupported conversion")
    }
}
