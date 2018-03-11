package host

import "testing"

func TestURLCheck(t *testing.T) {
	validURLs := []string{
		"http://www.foo.com",
		"ftp://www.foo.com",
		"http://www.foo.com/index.php?page=1",
	}

	invalidURLs := []string{
		"http://www.foo.kom",
		"htp://www.foo.com",
		"foo/bar",
		"foo",
		"www.foo.com",
		"http://vcruut.info",
		"http://puresafesupply.ru",
		"http://goo.gl/pero",
	}

	for _, validURL := range validURLs {
		err := IsValid(validURL)
		if err != nil {
			t.Error(validURL, "is not valid", err)
		}
	}

	for _, invalidURL := range invalidURLs {
		err := IsValid(invalidURL)
		if err == nil {
			t.Error(invalidURL, "is valid", err)
		}
	}
}
