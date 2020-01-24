## ADD THIS PART TO THE BEGINING OF $KAFKA_HOME/bin/kafka-run-class.sh
ENVVARSFILE="/Users/tcoueffe/path/to/envvars.sh"
if [ -f "${ENVVARSFILE}" ]; then
  source "${ENVVARSFILE}"
  if [ $? -ne 0 ]; then
    echo "Errors while loading ../config/envvars.sh file"
    exit 3
  else
    echo "${ENVVARSFILE} was sourced successfully !"
  fi
fi
