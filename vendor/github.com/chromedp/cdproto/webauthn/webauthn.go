// Package webauthn provides the Chrome DevTools Protocol
// commands, types, and events for the WebAuthn domain.
//
// This domain allows configuring virtual authenticators to test the WebAuthn
// API.
//
// Generated by the cdproto-gen command.
package webauthn

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// EnableParams enable the WebAuthn domain and start intercepting credential
// storage and retrieval with a virtual authenticator.
type EnableParams struct{}

// Enable enable the WebAuthn domain and start intercepting credential
// storage and retrieval with a virtual authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes WebAuthn.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// DisableParams disable the WebAuthn domain.
type DisableParams struct{}

// Disable disable the WebAuthn domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes WebAuthn.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// AddVirtualAuthenticatorParams creates and adds a virtual authenticator.
type AddVirtualAuthenticatorParams struct {
	Options *VirtualAuthenticatorOptions `json:"options"`
}

// AddVirtualAuthenticator creates and adds a virtual authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-addVirtualAuthenticator
//
// parameters:
//   options
func AddVirtualAuthenticator(options *VirtualAuthenticatorOptions) *AddVirtualAuthenticatorParams {
	return &AddVirtualAuthenticatorParams{
		Options: options,
	}
}

// AddVirtualAuthenticatorReturns return values.
type AddVirtualAuthenticatorReturns struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId,omitempty"`
}

// Do executes WebAuthn.addVirtualAuthenticator against the provided context.
//
// returns:
//   authenticatorID
func (p *AddVirtualAuthenticatorParams) Do(ctx context.Context) (authenticatorID AuthenticatorID, err error) {
	// execute
	var res AddVirtualAuthenticatorReturns
	err = cdp.Execute(ctx, CommandAddVirtualAuthenticator, p, &res)
	if err != nil {
		return "", err
	}

	return res.AuthenticatorID, nil
}

// RemoveVirtualAuthenticatorParams removes the given authenticator.
type RemoveVirtualAuthenticatorParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// RemoveVirtualAuthenticator removes the given authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-removeVirtualAuthenticator
//
// parameters:
//   authenticatorID
func RemoveVirtualAuthenticator(authenticatorID AuthenticatorID) *RemoveVirtualAuthenticatorParams {
	return &RemoveVirtualAuthenticatorParams{
		AuthenticatorID: authenticatorID,
	}
}

// Do executes WebAuthn.removeVirtualAuthenticator against the provided context.
func (p *RemoveVirtualAuthenticatorParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandRemoveVirtualAuthenticator, p, nil)
}

// AddCredentialParams adds the credential to the specified authenticator.
type AddCredentialParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}

// AddCredential adds the credential to the specified authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-addCredential
//
// parameters:
//   authenticatorID
//   credential
func AddCredential(authenticatorID AuthenticatorID, credential *Credential) *AddCredentialParams {
	return &AddCredentialParams{
		AuthenticatorID: authenticatorID,
		Credential:      credential,
	}
}

// Do executes WebAuthn.addCredential against the provided context.
func (p *AddCredentialParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandAddCredential, p, nil)
}

// GetCredentialParams returns a single credential stored in the given
// virtual authenticator that matches the credential ID.
type GetCredentialParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	CredentialID    string          `json:"credentialId"`
}

// GetCredential returns a single credential stored in the given virtual
// authenticator that matches the credential ID.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-getCredential
//
// parameters:
//   authenticatorID
//   credentialID
func GetCredential(authenticatorID AuthenticatorID, credentialID string) *GetCredentialParams {
	return &GetCredentialParams{
		AuthenticatorID: authenticatorID,
		CredentialID:    credentialID,
	}
}

// GetCredentialReturns return values.
type GetCredentialReturns struct {
	Credential *Credential `json:"credential,omitempty"`
}

// Do executes WebAuthn.getCredential against the provided context.
//
// returns:
//   credential
func (p *GetCredentialParams) Do(ctx context.Context) (credential *Credential, err error) {
	// execute
	var res GetCredentialReturns
	err = cdp.Execute(ctx, CommandGetCredential, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Credential, nil
}

// GetCredentialsParams returns all the credentials stored in the given
// virtual authenticator.
type GetCredentialsParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// GetCredentials returns all the credentials stored in the given virtual
// authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-getCredentials
//
// parameters:
//   authenticatorID
func GetCredentials(authenticatorID AuthenticatorID) *GetCredentialsParams {
	return &GetCredentialsParams{
		AuthenticatorID: authenticatorID,
	}
}

// GetCredentialsReturns return values.
type GetCredentialsReturns struct {
	Credentials []*Credential `json:"credentials,omitempty"`
}

// Do executes WebAuthn.getCredentials against the provided context.
//
// returns:
//   credentials
func (p *GetCredentialsParams) Do(ctx context.Context) (credentials []*Credential, err error) {
	// execute
	var res GetCredentialsReturns
	err = cdp.Execute(ctx, CommandGetCredentials, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Credentials, nil
}

// ClearCredentialsParams clears all the credentials from the specified
// device.
type ClearCredentialsParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
}

// ClearCredentials clears all the credentials from the specified device.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-clearCredentials
//
// parameters:
//   authenticatorID
func ClearCredentials(authenticatorID AuthenticatorID) *ClearCredentialsParams {
	return &ClearCredentialsParams{
		AuthenticatorID: authenticatorID,
	}
}

// Do executes WebAuthn.clearCredentials against the provided context.
func (p *ClearCredentialsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearCredentials, p, nil)
}

// SetUserVerifiedParams sets whether User Verification succeeds or fails for
// an authenticator. The default is true.
type SetUserVerifiedParams struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	IsUserVerified  bool            `json:"isUserVerified"`
}

// SetUserVerified sets whether User Verification succeeds or fails for an
// authenticator. The default is true.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#method-setUserVerified
//
// parameters:
//   authenticatorID
//   isUserVerified
func SetUserVerified(authenticatorID AuthenticatorID, isUserVerified bool) *SetUserVerifiedParams {
	return &SetUserVerifiedParams{
		AuthenticatorID: authenticatorID,
		IsUserVerified:  isUserVerified,
	}
}

// Do executes WebAuthn.setUserVerified against the provided context.
func (p *SetUserVerifiedParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetUserVerified, p, nil)
}

// Command names.
const (
	CommandEnable                     = "WebAuthn.enable"
	CommandDisable                    = "WebAuthn.disable"
	CommandAddVirtualAuthenticator    = "WebAuthn.addVirtualAuthenticator"
	CommandRemoveVirtualAuthenticator = "WebAuthn.removeVirtualAuthenticator"
	CommandAddCredential              = "WebAuthn.addCredential"
	CommandGetCredential              = "WebAuthn.getCredential"
	CommandGetCredentials             = "WebAuthn.getCredentials"
	CommandClearCredentials           = "WebAuthn.clearCredentials"
	CommandSetUserVerified            = "WebAuthn.setUserVerified"
)
