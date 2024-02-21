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
      
    printf '[-n|--name <arg>]' 
	printf '[-h|--help]\n'
}

parse_commandline()
{
	while test $# -gt 0
	do
		_key="$1"
		case "$_key" in
              
			-n|--name)
				_arg_name="$2"
				shift
				;;
			--name=*)
				_arg_name="${_key##--name=}"
				;;
			-o*)
				_arg_name="${_key##-n}"
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
   
	if [ -z "$_arg_name" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'name'" 1
    fi
   
}


# Variable Initialization
 
_arg_name=


parse_commandline "$@"
validate_args

printf '===== ENV VARIABLES =======\n'
  
printf '===========================\n'

printf '===== INPUT ARGUMENTS =====\n'
  
printf "  --name: "
maskPrint "$_arg_name"
  
printf '===========================\n'

 echo "hello "$_arg_name" !!" 

