OS := $(shell uname)
ARCH := $(shell uname -m)

ensure_tailwind_cli:
	@if [ ! -e tailwindcss ]; then\
		echo "No installation found";\
		if [ $(OS) = 'Darwin' ] && [ $(ARCH) = 'arm64' ]; then\
			echo "Beginning installation";\
			curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64;\
			chmod +x tailwindcss-macos-arm64;\
			mv tailwindcss-macos-arm64 tailwindcss;\
			echo "Tailwind is now installed as 'tailwindcss'";\
        else\
          echo "Your OS or arch is not supported\nPlease manually install tailwind CLI as 'tailwindcss' in this directory";\
        fi;\
	else\
		echo "Tailwind is already installed.";\
	fi;

tailwind-watch-templates:
	@./tailwindcss --watch -o ./resources/public/css/template-styles.css

tailwind-watch-components:
	@./tailwindcss --watch -i ./resources/css/components.css -o ./resources/public/css/template-styles.css
