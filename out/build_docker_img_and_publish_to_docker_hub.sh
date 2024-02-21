#!/bin/bash

# Created by argbash-init v2.10.0


# ARG_OPTIONAL_SINGLE([username],[o])

# ARG_OPTIONAL_SINGLE([password])

# ARG_OPTIONAL_SINGLE([path])

# ARG_OPTIONAL_SINGLE([image-name])

# ARG_OPTIONAL_SINGLE([image-tag])


# ARG_HELP([<The general help message of my script>])
# ARGBASH_GO()
# needed because of Argbash --> m4_ignore([
### START OF CODE GENERATED BY Argbash v2.10.0 one line above ###
# Argbash is a bash code generator used to get arguments parsing right.
# Argbash is FREE SOFTWARE, see https://argbash.io for more info


die()
{
	local _ret="${2:-1}"
	test "${_PRINT_HELP:-no}" = yes && print_help >&2
	echo "$1" >&2
	exit "${_ret}"
}


begins_with_short_option()
{
	local first_option all_short_options='oh'
	first_option="${1:0:1}"
	test "$all_short_options" = "${all_short_options/$first_option/}" && return 1 || return 0
}

# THE DEFAULTS INITIALIZATION - OPTIONALS
_arg_username=
_arg_password=
_arg_path=
_arg_image_name=
_arg_image_tag=


print_help()
{
	printf '%s\n' "<The general help message of my script>"
	printf 'Usage: %s [-o|--username <arg>] [--password <arg>] [--path <arg>] [--image-name <arg>] [--image-tag <arg>] [-h|--help]\n' "$0"
	printf '\t%s\n' "-h, --help: Prints help"
}


parse_commandline()
{
	while test $# -gt 0
	do
		_key="$1"
		case "$_key" in
			-o|--username)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
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
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_password="$2"
				shift
				;;
			--password=*)
				_arg_password="${_key##--password=}"
				;;
			--path)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_path="$2"
				shift
				;;
			--path=*)
				_arg_path="${_key##--path=}"
				;;
			--image-name)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_image_name="$2"
				shift
				;;
			--image-name=*)
				_arg_image_name="${_key##--image-name=}"
				;;
			--image-tag)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
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

parse_commandline "$@"

# OTHER STUFF GENERATED BY Argbash

### END OF CODE GENERATED BY Argbash (sortof) ### ])
# [ <-- needed because of Argbash

#

maskPrint () {
  local perc=75  ## percent to obfuscate
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

printf '===== ENV VARIABLES ====='

printf '%s: ' 'TEST_ENV'
maskPrint "$TEST_ENV"


printf '===== INPUT ARGUMENTS ====='

printf 'Value of --%s: %s\n' 'username' "$_arg_username"

printf 'Value of --%s: %s\n' 'password' "$_arg_password"

printf 'Value of --%s: %s\n' 'path' "$_arg_path"

printf 'Value of --%s: %s\n' 'image-name' "$_arg_image_name"

printf 'Value of --%s: %s\n' 'image-tag' "$_arg_image_tag"


docker image build -t "$_arg_image_name":"$_arg_tag" -f $_arg_path .
docker login -u "$_arg_username" -p "$_arg_password"
docker tag "$_arg_image_name" "$_arg_username"/"$_arg_image_name"
docker image push "$_arg_username"/"$_arg_image_name"


#
# ] <-- needed because of Argbash
