# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empirical_distribution.pdf

empirical_distribution.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empirical_distribution.tex

empirical_distribution.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_distribution.tex
	# pdflatex --file-line-error -interaction=nonstopmode empirical_distribution.tex
	../../latexrun empirical_distribution.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_distribution.pdf
	open empirical_distribution.pdf

o:
	make open

terms: empirical_distribution.tex
	bbk_sheet -mode ts empirical_distribution.tex

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
	rm -f empirical_distribution.*
	make reset
	bbk_sheet -mode mk > Makefile