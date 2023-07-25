#! /bin/bash
# Script used to chage the DEVELOPMENT variable between true and false.

if [ -f ./.env ]
then
    set -o allexport; source .env; set +o allexport
    printf "Environment loaded\n"
else
    printf "Please initialize the project, use: \033[1mmake init\033[0m\n"
    exit 0
fi

if [[ -z "${DEVELOPMENT}" ]]
then
    printf "Environment doesn't have DEVELOPMENT variable.\n"

    development=""
    developmentAnswer=true
    while $developmentAnswer; do
        echo -e "\n"
        read -r -p "Running development? (Y/N): " answer
        case $answer in
            [Yy]* )
                development="true"
                developmentAnswer=false
                ;;
            [Nn]* )
                development="false"
                developmentAnswer=false
                ;;
            * )
                printf "\nPlease answer Y or N."
                ;;
        esac
    done

    printf "\nDEVELOPMENT=\"$development\"\n" >> ./.env
else
    if [[ "${DEVELOPMENT}" == "true" ]]
    then
        printf "Environment changed to production\n"
        echo "$(awk -F "=" '/^DEVELOPMENT/ {gsub("true", "false")}1' .env)" > .env
    else
        printf "Environment changed to development\n"
        echo "$(awk -F "=" '/^DEVELOPMENT/ {gsub("false", "true")}1' .env)" > .env
    fi
fi
