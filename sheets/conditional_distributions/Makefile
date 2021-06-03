# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make conditional_distributions.pdf

conditional_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > conditional_distributions.tex

conditional_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_distributions.tex
	# pdflatex --file-line-error -interaction=nonstopmode conditional_distributions.tex
	../../latexrun conditional_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: conditional_distributions.pdf
	open conditional_distributions.pdf

o:
	make open

terms: conditional_distributions.tex
	bbk_sheet -mode ts conditional_distributions.tex

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
	rm -f conditional_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile