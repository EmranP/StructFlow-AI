package error

import "errors"

var (
	ErrUserAlreadyExists = errors.New(
		"user already exists",
	)

	ErrUserNotFound       = errors.New("user not found")
	ErrProjectNotFound    = errors.New("project not found")
	ErrGenerationNotFound = errors.New(
		"generation not found",
	)
	ErrEmailNotVerified = errors.New(
		"email not verified",
	)
	ErrVerificationCodeResent = errors.New(
		"verification code resent",
	)
	ErrVerificationCooldown = errors.New(
		"please wait before requesting a new verification code",
	)
	ErrSessionNotFound = errors.New(
		"Session not found!",
	)

	ErrAiProviderNotFound = errors.New(
		"AI provider not found",
	)
	ErrAiProviderInvalid            = errors.New("invalid model")
	ErrAiProviderGeminiUnavailable  = errors.New("gemini unavailable")
	ErrAiProviderClaudeUnavailable  = errors.New("claude unavailable")
	ErrAiProviderChatGPTUnavailable = errors.New("chatGPT unavailable")

	ErrInvalidVerificationCode = errors.New(
		"invalid verification code",
	)
	ErrInvalidCredentials = errors.New(
		"invalid credentials",
	)

	ErrUnauthorized = errors.New(
		"unauthorized",
	)
)
