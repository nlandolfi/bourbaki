for d in */; do
  echo "$d"
  cd $d && bbk_sheet >> Makefile
done

