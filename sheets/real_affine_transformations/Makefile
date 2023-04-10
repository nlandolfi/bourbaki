# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_affine_transformations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_affine_transformations.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > real_affine_transformations.tex

real_affine_transformations.pdf: ../../*.tex ../../trademark.pdf *.tex real_affine_transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_affine_transformations.tex
	../../latexrun real_affine_transformations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_affine_transformations.pdf
	open real_affine_transformations.pdf

o:
	make open

terms: real_affine_transformations.tex
	bbk_sheet -mode ts real_affine_transformations.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_affine_transformations.*
	make reset
	bbk_sheet -mode mk > Makefile