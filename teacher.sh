#! /bin/bash
INTMURDER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2)
echo "$INTMURDER"
cat ./interviews/interview-$INTMURDER
echo "$MAIN_SUSPECT"