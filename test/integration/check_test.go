package integration

import (
	"os"
	"testing"

	"yadro/internal/pkg"
)

func TestCheckApp(t *testing.T) {
	tests := []struct {
		ID         string
		checkFiles string
		result     string
		err        error
	}{
		{
			ID:         "Test 1",
			checkFiles: "testFiles\\test_file_1.txt",
			result:     "09:00\r\n08:48 1 client1\r\n08:48 13 NotOpenYet\r\n09:41 1 client1\r\n09:48 1 client2\r\n09:52 3 client1\r\n09:52 13 ICanWaitNoLonger!\r\n09:54 2 client1 1\r\n10:25 2 client2 2\r\n10:58 1 client3\r\n10:59 2 client3 3\r\n11:30 1 client4\r\n11:35 2 client4 2\r\n11:35 13 PlaceIsBusy\r\n11:45 3 client4\r\n12:33 4 client1\r\n12:33 12 client4 1\r\n12:43 4 client2\r\n15:52 4 client4\r\n19:00 11 client3\r\n19:00\r\n1 70 05:58\r\n2 30 02:18\r\n3 90 08:01",
		},
		{
			ID:         "Test 2",
			checkFiles: "testFiles\\test_file_2.txt",
			result:     "09:00\r\n09:00 1 client1\r\n09:00 2 client1 1\r\n09:00 1 client2\r\n09:00 2 client2 2\r\n10:00 1 client3\r\n10:00 3 client3\r\n10:00 1 client4\r\n10:00 3 client4\r\n10:00 1 client5\r\n10:00 3 client5\r\n10:00 11 client5\r\n11:00 4 client1\r\n11:00 12 client3 1\r\n12:00 4 client2\r\n12:00 12 client4 2\r\n13:00 4 client3\r\n13:00 4 client4\r\n19:00\r\n1 40 04:00\r\n2 40 04:00",
		},
		{
			ID:         "Test 3",
			checkFiles: "testFiles\\test_file_3.txt",
			err:        pkg.InvalidNumberOfLines,
		},
		{
			ID:         "Test 4",
			checkFiles: "testFiles\\test_file_4.txt",
			err:        pkg.BadFormatOfLine,
			result:     "09:00 1 client?*%2",
		},
	}

	for _, test := range tests {
		t.Run(test.ID, func(t *testing.T) {
			os.Args = append([]string{"program_name"}, test.checkFiles)
			actual, err := pkg.Parse(test.checkFiles)
			if err != nil {
				if err != test.err {
					t.Errorf("Expected : %s, actual: %s", test.err, err)
				}
			}

			if actual != test.result {
				t.Errorf("Expected: %s, actual: %s", test.result, actual)
			}
		})
	}
}
