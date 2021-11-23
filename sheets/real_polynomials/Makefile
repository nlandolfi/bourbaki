# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_polynomials.pdf

real_polynomials.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_polynomials.tex

real_polynomials.pdf: ../../*.tex ../../trademark.pdf *.tex real_polynomials.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_polynomials.tex
	../../latexrun real_polynomials.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_polynomials.pdf
	open real_polynomials.pdf

o:
	make open

terms: real_polynomials.tex
	bbk_sheet -mode ts real_polynomials.tex

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
	rm -f real_polynomials.*
	make reset
	bbk_sheet -mode mk > Makefile