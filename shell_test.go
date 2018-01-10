package tlkn

import (
	"context"
	"testing"
)

func TestBash(t *testing.T) {
	type args struct {
		cmd string
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
		wantErr bool
	}{
		{"basic", args{"echo \"Test\"", context.Background()}, "Test\n", false},
		{"error-basic", args{"exit 1", context.Background()}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Bash(tt.args.ctx, tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bash() error = %v, wantErr %#v", err, tt.wantErr)
			}
			if string(out) != tt.wantRes {
				t.Errorf("Bash() error = %v, wantRes= %#v gotRes= %#v", err, tt.wantRes, string(out))
			}
		})
	}
}

func TestPrompt(t *testing.T) {
	type args struct {
		prompt string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prompt(tt.args.prompt, tt.args.args...); got != tt.want {
				t.Errorf("Prompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromptRequired(t *testing.T) {
	type args struct {
		prompt string
		args   []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := PromptRequired(tt.args.prompt, tt.args.args...); gotS != tt.wantS {
				t.Errorf("PromptRequired() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestPromptConfirm(t *testing.T) {
	type args struct {
		prompt string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PromptConfirm(tt.args.prompt, tt.args.args...); got != tt.want {
				t.Errorf("PromptConfirm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromptChoose(t *testing.T) {
	type args struct {
		prompt string
		list   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PromptChoose(tt.args.prompt, tt.args.list); got != tt.want {
				t.Errorf("PromptChoose() = %v, want %v", got, tt.want)
			}
		})
	}
}
