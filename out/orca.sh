#!/bin/bash

die()
{
	local _ret="${2:-1}"
	test "${_PRINT_HELP:-no}" = yes && print_help >&2
	echo "$1" >&2
	exit "${_ret}"
}

maskPrint () {
  local perc=100  ## percent to obfuscate
  local i=0
  for((i=0; i < ${#1}; i++))
  do
    if [ $(( $RANDOM % 100 )) -lt "$perc" ]
    then
        printf '%s' '*'
    else
        printf '%s' "${1:i:1}"
    fi
  done
  echo
}

print_help()
{
	printf '%s\n' "<The general help message of my script>"
	printf 'Usage: %s ' "$0"
       
    printf '[-o|--username <arg>]'   
    printf '[--password <arg>]'   
    printf '[--path <arg>]'   
    printf '[--image-name <arg>]'   
    printf '[--image-tag <arg>]' 
	printf '[-h|--help]\n'
}

parse_commandline()
{
	while test $# -gt 0
	do
		_key="$1"
		case "$_key" in
               
			-o|--username)
				_arg_username="$2"
				shift
				;;
			--username=*)
				_arg_username="${_key##--username=}"
				;;
			-o*)
				_arg_username="${_key##-o}"
				;;
  
			--password)
				_arg_password="$2"
				shift
				;;
			--password=*)
				_arg_password="${_key##--password=}"
				;;  
			--path)
				_arg_path="$2"
				shift
				;;
			--path=*)
				_arg_path="${_key##--path=}"
				;;  
			--image-name)
				_arg_image_name="$2"
				shift
				;;
			--image-name=*)
				_arg_image_name="${_key##--image-name=}"
				;;  
			--image-tag)
				_arg_image_tag="$2"
				shift
				;;
			--image-tag=*)
				_arg_image_tag="${_key##--image-tag=}"
				;;
			-h|--help)
				print_help
				exit 0
				;;
			-h*)
				print_help
				exit 0
				;;
			*)
				_PRINT_HELP=yes die "FATAL ERROR: Got an unexpected argument '$1'" 1
				;;
		esac
		shift
	done
}

validate_args() {
     
	if [ -z "$_arg_username" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'username'" 1
    fi
    
	if [ -z "$_arg_password" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'password'" 1
    fi
    
	if [ -z "$_arg_path" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'path'" 1
    fi
    
	if [ -z "$_arg_image_name" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'image-name'" 1
    fi
    
	if [ -z "$_arg_image_tag" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'image-tag'" 1
    fi
   
}


# Variable Initialization
  
_arg_username=
 
_arg_password=
 
_arg_path=
 
_arg_image_name=
 
_arg_image_tag=


parse_commandline "$@"
validate_args

printf '===== ENV VARIABLES =======\n'
  
printf "  TEST_ENV: %s\n" "$TEST_ENV"
            
printf '===========================\n'

printf '===== INPUT ARGUMENTS =====\n'
    
printf "  --username: "
maskPrint "$_arg_username"
    
printf "  --password: %s\n" "$_arg_password"
    
printf "  --path: %s\n" "$_arg_path"
    
printf "  --image-name: %s\n" "$_arg_image_name"
    
printf "  --image-tag: %s\n" "$_arg_image_tag"
  
printf '===========================\n'

echo "$_arg_username"
echo "$_arg_password"
echo "$_arg_path"
echo "$_arg_image_name"
echo "$_arg_image_tag"

