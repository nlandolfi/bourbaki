# generated automatically by sheet.go

make:
	make closest_point_property.pdf

closest_point_property.tex: sheet.tex
	sheet -mode c -in sheet.tex > closest_point_property.tex

closest_point_property.pdf: ../../*.tex ../../trademark.pdf *.tex closest_point_property.tex
	pdflatex closest_point_property.tex
	make defs


open: closest_point_property.pdf
	open closest_point_property.pdf

o:
	make open

defs: closest_point_property.tex
	defs closest_point_property.tex

spell:
	aspell -c closest_point_property.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile