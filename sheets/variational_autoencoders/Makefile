# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make variational_autoencoders.pdf

variational_autoencoders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > variational_autoencoders.tex

variational_autoencoders.pdf: ../../*.tex ../../trademark.pdf *.tex variational_autoencoders.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode variational_autoencoders.tex
	../../latexrun variational_autoencoders.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: variational_autoencoders.pdf
	open variational_autoencoders.pdf

o:
	make open

terms: variational_autoencoders.tex
	bbk_sheet -mode ts variational_autoencoders.tex

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
	rm -f variational_autoencoders.*
	make reset
	bbk_sheet -mode mk > Makefile