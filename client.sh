#!/bin/bash

base_url="http://localhost:8180"

list_messages() {
    curl -s "${base_url}/message"
}

create_message() {
	auth_token="$1"
    content="$2"
    curl -s -X POST "${base_url}/message" -H "Content-Type: application/json" -H "client_id: ${auth_token}" -d "{\"content\": \"${content}\"}"
}

update_message() {
	auth_token="$1"
    client_id="$2"
    content="$3"
    curl -s -X PUT "${base_url}/message/${client_id}" -H "Content-Type: application/json" -H "client_id: ${auth_token}" -d "{\"content\": \"${content}\"}"
}

get_message() {
    client_id="$1"
    curl -s "${base_url}/message/${client_id}"
}

delete_message() {
	auth_token="$1"
    client_id="$2"
    curl -s -X DELETE "${base_url}/message/${client_id}" -H "client_id: ${auth_token}"
}

if [ $# -lt 1 ]; then
    echo "Usage: $0 <action>"
    exit 1
fi

action="$1"
case "$action" in
    "list")
        list_messages
        ;;
    "create")
        if [ $# -lt 3 ]; then
            echo "Usage: $0 create_message <auth_id> <content>"
            exit 1
        fi
        create_message "$2" "$3"
        ;;
    "update")
        if [ $# -lt 4 ]; then
            echo "Usage: $0 update_message <auth_id> <client_id_to_update> <content>"
            exit 1
        fi
        update_message "$2" "$3" "$4"
        ;;
    "get")
        if [ $# -lt 2 ]; then
            echo "Usage: $0 get_message <client_id>"
            exit 1
        fi
        get_message "$2"
        ;;
    "delete")
        if [ $# -lt 3 ]; then
            echo "Usage: $0 delete_message <auth_id> <client_id_to_delete>"
            exit 1
        fi
        delete_message "$2" "$3"
        ;;
    *)
        echo "Invalid action: $action"
        exit 1
        ;;
esac
