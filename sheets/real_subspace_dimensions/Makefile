# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_subspace_dimensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_subspace_dimensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_subspace_dimensions.tex

real_subspace_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex real_subspace_dimensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_subspace_dimensions.tex
	../../latexrun real_subspace_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_subspace_dimensions.pdf
	open real_subspace_dimensions.pdf

o:
	make open

terms: real_subspace_dimensions.tex
	bbk_sheet -mode ts real_subspace_dimensions.tex

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
	rm -f real_subspace_dimensions.*
	make reset
	bbk_sheet -mode mk > Makefile