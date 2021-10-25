# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make state_representations.pdf

state_representations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > state_representations.tex

state_representations.pdf: ../../*.tex ../../trademark.pdf *.tex state_representations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode state_representations.tex
	../../latexrun state_representations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: state_representations.pdf
	open state_representations.pdf

o:
	make open

terms: state_representations.tex
	bbk_sheet -mode ts state_representations.tex

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
	rm -f state_representations.*
	make reset
	bbk_sheet -mode mk > Makefile