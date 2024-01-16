#!/bin/bash
SCALING_FACTOR=10

# Run pgbench to generate and load mock data
PGPASSWORD=va3dvcm0xNA pgbench -iq -s ${SCALING_FACTOR} -U bookworm products