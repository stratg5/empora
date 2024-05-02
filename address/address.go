package address

import (
	"empora/entities"
	"fmt"
	"strings"

	street "github.com/smartystreets/smartystreets-go-sdk/us-street-api"
)

type Service struct {
	LookupClient lookupSender
}

// the address service handles any address related operations including sending and formatting
func NewService(lookupSender lookupSender) Service {
	return Service{
		LookupClient: lookupSender,
	}
}

func (s Service) BuildLookupsFromAddresses(addresses []entities.Address) []*street.Lookup {
	lookups := []*street.Lookup{}
	for _, address := range addresses {
		lookups = append(lookups, &street.Lookup{
			Street:  address.Street,
			City:    address.Street,
			ZIPCode: address.ZipCode,
		})
	}

	return lookups
}

func (s Service) BuildRawDataFromLookups(addresses []entities.Address, lookups []*street.Lookup) [][]string {
	output := [][]string{}
	for idx, lookup := range lookups {
		var cols []string
		tempAddr := addresses[idx]

		if !tempAddr.Valid || len(lookup.Results) == 0 {
			// TODO split the lastline to add comma?
			cols = []string{tempAddr.Street, " " +tempAddr.City, tempAddr.ZipCode + " -> Invalid Address"}
		} else {
			lookupAddr := lookup.Results[0]
			// TODO split the lastline to add comma?
			cols = []string{tempAddr.Street, " " + tempAddr.City, " " + tempAddr.ZipCode + " -> " + lookupAddr.DeliveryLine1, " " + lookupAddr.LastLine}
		}

		output = append(output, cols)
	}

	return output
}

// BuildAddresses takes in the raw CSV data and builds an address array
func (s Service) BuildAddressesFromRawData(data [][]string) []entities.Address {
	addresses := []entities.Address{}
	for idx, row := range data {
		// skip the column names
		if idx == 0 {
			continue
		}

		// check if the row is an invalid length
		if len(row) < 3 || len(row) > 3 {
			originString := ""
			for _, col := range row {
				originString += col
			}

			address := entities.Address{
				OriginString: originString,
				Valid:        false,
			}

			addresses = append(addresses, address)
			continue
		}

		// trim the leading and trailing spaces from the data
		address := entities.Address{
			Street:  strings.TrimSpace(row[0]),
			City:    strings.TrimSpace(row[1]),
			ZipCode: strings.TrimSpace(row[2]),
			Valid:   true,
		}

		addresses = append(addresses, address)
	}
	return addresses
}

func (s Service) SendLookups(lookups ...*street.Lookup) error {
	err := s.LookupClient.SendLookups(lookups...)
	if err != nil {
		return fmt.Errorf("error while sending lookups: %w", err)
	}

	return nil
}
