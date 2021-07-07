# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make maximum_likelihood.pdf

maximum_likelihood.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > maximum_likelihood.tex

maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex maximum_likelihood.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode maximum_likelihood.tex
	../../latexrun maximum_likelihood.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: maximum_likelihood.pdf
	open maximum_likelihood.pdf

o:
	make open

terms: maximum_likelihood.tex
	bbk_sheet -mode ts maximum_likelihood.tex

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
	rm -f maximum_likelihood.*
	make reset
	bbk_sheet -mode mk > Makefile