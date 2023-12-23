package utils

import (
	"strconv"
)

func Atoi(out *int64, src string) {
	result, err := strconv.Atoi(src)
	if err != nil {
		// var id = uuid.New()
		// log.Warn().Msgf("convert to str int, use default value of int. id=%s err=%v", id, err)
		if *out == 0 {
			*out = 0
		}
		return
	}

	*out = int64(result)
}
