package main

import (
	"testing"

	"github.com/pikopod/pikopod/internal/config"
)

func TestNewLLMClientBaseURLOverrides(t *testing.T) {
	cfg := &config.Config{LLM: config.LLM{
		Provider:      "openrouter",
		APIKey:        "test-key",
		OpenRouterKey: "test-key",
	}}

	t.Setenv("PIKOPOD_LLM_BASE", "https://generic.example/v1")
	t.Setenv("PIKOPOD_OPENROUTER_BASE", "https://legacy.example/v1")
	if got := newLLMClient(cfg, "").BaseURL; got != "https://generic.example/v1" {
		t.Fatalf("generic base override = %q", got)
	}

	t.Setenv("PIKOPOD_LLM_BASE", "")
	if got := newLLMClient(cfg, "").BaseURL; got != "https://legacy.example/v1" {
		t.Fatalf("legacy OpenRouter base override = %q", got)
	}
}
