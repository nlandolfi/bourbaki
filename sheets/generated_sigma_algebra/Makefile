# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make generated_sigma_algebra.pdf
	make graph.pdf

generated_sigma_algebra.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > generated_sigma_algebra.tex

generated_sigma_algebra.pdf: ../../*.tex ../../trademark.pdf *.tex generated_sigma_algebra.tex
	pdflatex --file-line-error -interaction=nonstopmode generated_sigma_algebra.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generated_sigma_algebra.pdf
	open generated_sigma_algebra.pdf

o:
	make open

defs: generated_sigma_algebra.tex
	defs generated_sigma_algebra.tex

spell:
	aspell -c generated_sigma_algebra.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f generated_sigma_algebra.*
	make reset
	bbk_sheet -mode mk > Makefile