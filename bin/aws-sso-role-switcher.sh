#!/bin/sh
if [[ $* == *-h* ]]; then
  aws-sso-role-switcher $@
elif [[ $* == *-v* ]]; then
    aws-sso-role-switcher $@
else
  eval "$(aws-sso-role-switcher "$@")"
fi