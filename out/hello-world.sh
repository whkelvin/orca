#!/bin/bash
set -e

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
    printf '[-a|--adj <arg>]'  
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
  
			-a|--adj)
				_arg_adj="$2"
				shift
				;;
			--adj=*)
				_arg_adj="${_key##--adj=}"
				;;
			-o*)
				_arg_adj="${_key##-a}"
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
        
	if [ -z "$TEST_ENV_VAR" ]
    then
      _PRINT_HELP=yes die "Missing ENV Variable: 'TEST_ENV_VAR'" 1
    fi
    

        
	if [ -z "$_arg_adj" ]
    then
      _PRINT_HELP=yes die "Missing value for arg: 'adj'" 1
    fi
      
}


# Variable Initialization
 
_arg_name=
 
_arg_adj=
 

parse_commandline "$@"
validate_args

printf '===== ENV VARIABLES =======\n'
      
printf "  TEST_ENV_VAR: "
maskPrint "$TEST_ENV_VAR"
  
printf '===========================\n'

printf '===== INPUT ARGUMENTS =====\n'
  
printf "  --name: %s\n" "$_arg_name"
    
printf "  --adj: %s\n" "$_arg_adj"
    
printf '===========================\n'

echo "hello "$_arg_name", you are "$_arg_adj" !!"
echo "TEST_ENV_VAR is "$TEST_ENV_VAR""

