# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make empirical_measure.pdf
	make graph.pdf

empirical_measure.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empirical_measure.tex

empirical_measure.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_measure.tex
	# pdflatex --file-line-error -interaction=nonstopmode empirical_measure.tex
	../../latexrun empirical_measure.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_measure.pdf
	open empirical_measure.pdf

o:
	make open

terms: empirical_measure.tex
	bbk_sheet -mode ts empirical_measure.tex

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
	rm -f empirical_measure.*
	make reset
	bbk_sheet -mode mk > Makefile