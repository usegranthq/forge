SERVICE_NAME=backend
BINARY_NAME=api

REMOTE_APP_DIR=/home/ubuntu/apps/${SERVICE_NAME}/
REMOTE_DEPLOY_PATH=${REMOTE_APP_DIR}tmp/

# stop pm2 service
pm2 stop ${SERVICE_NAME}

# remove binary
rm -rf ${REMOTE_APP_DIR}${BINARY_NAME}

# copy binary from tmp directory to app directory
cp ${REMOTE_DEPLOY_PATH}${BINARY_NAME} ${REMOTE_APP_DIR}${BINARY_NAME}

# start pm2 service
pm2 start ${SERVICE_NAME}
