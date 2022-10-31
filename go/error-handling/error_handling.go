package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()
	for errors.As(err, &TransientError{}) {
		resource, err = opener()
	}
	if err != nil {
		return err
	}
	defer resource.Close()

	defer func() {
		if recErr := recover(); recErr != nil {
			panicErr := recErr.(error)

			var frobErr FrobError
			if errors.As(panicErr, &frobErr) {
				resource.Defrob(frobErr.defrobTag)
			}

			err = panicErr
		}
	}()

	resource.Frob(input)

	return nil
}
