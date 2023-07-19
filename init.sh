#! /bin/bash

# Not super efficient but it gets the job done.

# Variables to check and insert
SUPABASE_PROJECT_PASSWORD=""
SUPABASE_URL=""
SUPABASE_API_PUBLIC_KEY=""
SUPABASE_API_PRIVATE_KEY=""
SUPABASE_DB_HOST=""
SUPABASE_DB_PORT=""
SUPABASE_DB_NAME=""
SUPABASE_DB_USER=""
SUPABASE_DB_SSLMODE=""
SUPABASE_DB_SSLCERT_PATH=""

STRIPE_API_KEY=""

PORT="" # "" := 3000
DEVELOPMENT=""

# Create the .env file if not present
if [ -f ./.env ]
then
    echo ".env exists"
else
    touch .env
fi
echo -e "\n"

# Complete variables and add them to .env
read -r -p "SUPABASE_PROJECT_PASSWORD: " SUPABASE_PROJECT_PASSWORD
printf "SUPABASE_PROJECT_PASSWORD=\"$SUPABASE_PROJECT_PASSWORD\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_URL: " SUPABASE_URL
printf "SUPABASE_URL=\"$SUPABASE_URL\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_API_PUBLIC_KEY: " SUPABASE_API_PUBLIC_KEY
printf "SUPABASE_API_PUBLIC_KEY=\"$SUPABASE_API_PUBLIC_KEY\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_API_PRIVATE_KEY: " SUPABASE_API_PRIVATE_KEY
printf "SUPABASE_API_PRIVATE_KEY=\"$SUPABASE_API_PRIVATE_KEY\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_DB_HOST: " SUPABASE_DB_HOST
printf "SUPABASE_DB_HOST=\"$SUPABASE_DB_HOST\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_DB_PORT: " SUPABASE_DB_PORT
printf "SUPABASE_DB_PORT=\"$SUPABASE_DB_PORT\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_DB_NAME: " SUPABASE_DB_NAME
printf "SUPABASE_DB_NAME=\"$SUPABASE_DB_NAME\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_DB_USER: " SUPABASE_DB_USER
printf "SUPABASE_DB_USER=\"$SUPABASE_DB_USER\"\n" >> ./.env
echo -e "\n"

read -r -p "SUPABASE_DB_SSLMODE (disable/verify-full): " SUPABASE_DB_SSLMODE
printf "SUPABASE_DB_SSLMODE=\"$SUPABASE_DB_SSLMODE\"\n" >> ./.env
echo -e "\n"

if [[ "$SUPABASE_DB_SSLMODE" == "verify-full" ]]
then
    echo -e "\n"
    read -r -p "SUPABASE_DB_SSLCERT_PATH: (from app ROOT)" SUPABASE_DB_SSLCERT_PATH
    printf "SUPABASE_DB_SSLCERT_PATH=\"$SUPABASE_DB_SSLCERT_PATH\"\n\n" >> ./.env
else
    printf "SUPABASE_DB_SSLCERT_PATH=\"$SUPABASE_DB_SSLCERT_PATH\"\n\n" >> ./.env
fi


read -r -p "STRIPE_API_KEY: " STRIPE_API_KEY
printf "STRIPE_API_KEY=\"$STRIPE_API_KEY\"\n\n" >> ./.env


portAnswer=true
while $portAnswer; do
    echo -e "\n"
    read -r -p "Which port should I use? (No answer default to 3000): " PORT
    case $PORT in
        [0-9][0-9][0-9][0-9] )
            portAnswer=false
            ;;
        "" )
            portAnswer=false
            ;;
        * )
            printf "\nPlease answer a PORT or don't answer."
            ;;
    esac
done
printf "PORT=\"$PORT\"\n" >> ./.env

developmentAnswer=true
while $developmentAnswer; do
    echo -e "\n"
    read -r -p "Running development? (Y/N): " answer
    case $answer in
        [Yy]* )
            DEVELOPMENT="true"
            developmentAnswer=false
            ;;
        [Nn]* )
            DEVELOPMENT="false"
            developmentAnswer=false
            ;;
        * )
            printf "\nPlease answer Y or N."
            ;;
    esac
done
printf "DEVELOPMENT=\"$DEVELOPMENT\"\n" >> ./.env

echo -e "\n"
echo "Installing node packages"
npm i

echo -e "\n"
printf "Setup completed!\n\nTo start simply run: \033[1mmake\033[0m\n"
echo -e "\n"