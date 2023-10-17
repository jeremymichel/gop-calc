package gop_calc

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Calculate will return the GOP size in seconds and in frames based on the framerate, samplerate, audioframes
// and the target in seconds. audioFrames is typically 1024.
func Calculate(fps int, sampleRate int, audioFrames int, target int) (float64, int) {
	gopSize := fps * audioFrames / gcd(fps*audioFrames, sampleRate)
	gopFrames := gopSize
	gopSec := float64(gopFrames) / float64(fps)

	for gopSec <= float64(target) {
		if float64(gopFrames+gopSize)/float64(fps) > float64(target) {
			break
		}
		gopFrames += gopSize
		gopSec = float64(gopFrames) / float64(fps)
	}

	return gopSec, gopFrames
}
