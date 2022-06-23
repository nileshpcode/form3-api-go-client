package main

import (
	"encoding/json"
	"fmt"
	ferror "github.com/form3/error"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

// AccountInterface provides an interface to enable mocking the
type AccountInterface interface {
	Create(payload AccountRequestDTO) (*AccountResponseDTO, error)
	Fetch(id uuid.UUID) (*AccountResponseDTO, error)
	Delete(id uuid.UUID, version int) error
}

// AccountService provides the API operation methods for making requests to form3 Account
type accountService struct {
	client *resty.Client
}

func newAccountService(c *resty.Client) *accountService {
	return &accountService{c}
}

const (
	// AccountsAPIBaseURL indicates the base url of account api
	AccountsAPIBaseURL = "v1/organisation/accounts"
)

// Create creates an existing bank account with Form3 or create a new one. The country attribute must be specified as a minimum. Depending on the country, other attributes such as bank_id and bic are mandatory.
func (svc *accountService) Create(data AccountRequestDTO) (*AccountResponseDTO, error) {
	apiURL := fmt.Sprintf("%s/%s", svc.client.BaseURL, AccountsAPIBaseURL)

	resp, err := svc.client.R().SetBody(data).Post(apiURL)
	if err != nil {
		return nil, ferror.NewAPIClientError(apiURL, nil, nil, fmt.Errorf("unable to invoke API: %w", err))
	}

	defer resp.RawResponse.Body.Close()
	responseBody := string(resp.Body())
	if resp.StatusCode() != http.StatusCreated { // All 3xx, 4xx, 5xx are considered errors
		return nil, ferror.NewAPIClientError(apiURL, &resp.RawResponse.StatusCode, &responseBody, fmt.Errorf("received non-created code: %d", resp.StatusCode()))
	}

	responseDTO := AccountResponseDTO{}
	if err := json.NewDecoder(resp.RawResponse.Body).Decode(&responseDTO); err != nil {
		return nil, ferror.NewAPIClientError(apiURL, &resp.RawResponse.StatusCode, nil, fmt.Errorf("unable parse response payload: %w", err))
	}

	return &responseDTO, nil
}

// Fetch gets a single account using the account ID
func (svc *accountService) Fetch(id uuid.UUID) (*AccountResponseDTO, error) {
	apiURL := fmt.Sprintf("%s/%s/%v", svc.client.BaseURL, AccountsAPIBaseURL, id)

	resp, err := svc.client.R().Get(apiURL)
	if err != nil {
		fmt.Println(err, "Rsfewf")
		return nil, ferror.NewAPIClientError(apiURL, nil, nil, fmt.Errorf("unable to invoke API: %w", err))
	}
	fmt.Println("SFDf")
	defer resp.RawResponse.Body.Close()
	responseBody := string(resp.Body())
	if resp.StatusCode() != http.StatusOK { // All 3xx, 4xx, 5xx are considered errors
		return nil, ferror.NewAPIClientError(apiURL, &resp.RawResponse.StatusCode, &responseBody, fmt.Errorf("received non-ok code: %d", resp.StatusCode()))
	}

	responseDTO := AccountResponseDTO{}
	if err := json.NewDecoder(resp.RawResponse.Body).Decode(&responseDTO); err != nil {
		return nil, ferror.NewAPIClientError(apiURL, &resp.RawResponse.StatusCode, nil, fmt.Errorf("unable parse response payload: %w", err))
	}

	return &responseDTO, nil
}

// Delete deletes an account
func (svc *accountService) Delete(id uuid.UUID, version int) error {
	apiURL := fmt.Sprintf("%s/%s/%v", svc.client.BaseURL, AccountsAPIBaseURL, id)

	resp, err := svc.client.R().SetQueryParam("version", strconv.Itoa(version)).Delete(apiURL)
	if err != nil {
		return ferror.NewAPIClientError(apiURL, nil, nil, fmt.Errorf("unable to invoke API: %w", err))
	}

	defer resp.RawResponse.Body.Close()
	responseBody := string(resp.Body())
	if resp.StatusCode() != http.StatusOK { // All 3xx, 4xx, 5xx are considered errors
		return ferror.NewAPIClientError(apiURL, &resp.RawResponse.StatusCode, &responseBody, fmt.Errorf("received non-ok code: %d", resp.StatusCode()))
	}

	return nil
}

// AccountRequestDTO represents an account in the form3 org section.
type AccountRequestDTO struct {
	ID             uuid.UUID `json:"id,omitempty"`
	OrganisationID uuid.UUID `json:"organisation_id,omitempty"`
	Type           string    `json:"type,omitempty"`
	Version        *int64    `json:"version,omitempty"`
	Attributes     struct {
		AccountClassification   *string  `json:"account_classification,omitempty"`
		AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
		AccountNumber           string   `json:"account_number,omitempty"`
		AlternativeNames        []string `json:"alternative_names,omitempty"`
		BankID                  string   `json:"bank_id,omitempty"`
		BankIDCode              string   `json:"bank_id_code,omitempty"`
		BaseCurrency            string   `json:"base_currency,omitempty"`
		Bic                     string   `json:"bic,omitempty"`
		Country                 *string  `json:"country,omitempty"`
		Iban                    string   `json:"iban,omitempty"`
		JointAccount            *bool    `json:"joint_account,omitempty"`
		Name                    []string `json:"name,omitempty"`
		SecondaryIdentification string   `json:"secondary_identification,omitempty"`
		Status                  *string  `json:"status,omitempty"`
		Switched                *bool    `json:"switched,omitempty"`
	} `json:"attributes,omitempty"`
}

type AccountResponseDTO struct {
	Data struct {
		Type           string    `json:"type,omitempty"`
		ID             uuid.UUID `json:"id,omitempty"`
		Version        int       `json:"version,omitempty"`
		OrganisationID uuid.UUID `json:"organisation_id,omitempty"`
		Attributes     struct {
			Country         string `json:"country,omitempty"`
			BaseCurrency    string `json:"base_currency,omitempty"`
			AccountNumber   string `json:"account_number,omitempty"`
			BankID          string `json:"bank_id,omitempty"`
			BankIDCode      string `json:"bank_id_code,omitempty"`
			Bic             string `json:"bic,omitempty"`
			Iban            string `json:"iban,omitempty"`
			Status          string `json:"status,omitempty"`
			UserDefinedData []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"user_defined_data,omitempty"`
			ValidationType      string `json:"validation_type,omitempty"`
			ReferenceMask       string `json:"reference_mask,omitempty"`
			AcceptanceQualifier string `json:"acceptance_qualifier,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			AccountEvents struct {
				Data []struct {
					Type string    `json:"type,omitempty"`
					ID   uuid.UUID `json:"id,omitempty"`
				} `json:"data,omitempty"`
			} `json:"account_events,omitempty"`
		} `json:"relationships,omitempty"`
	} `json:"data,omitempty"`
}
