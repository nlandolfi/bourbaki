# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make empirical_normal.pdf
	make graph.pdf

empirical_normal.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empirical_normal.tex

empirical_normal.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_normal.tex
	# pdflatex --file-line-error -interaction=nonstopmode empirical_normal.tex
	../../latexrun empirical_normal.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_normal.pdf
	open empirical_normal.pdf

o:
	make open

terms: empirical_normal.tex
	bbk_sheet -mode ts empirical_normal.tex

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
	rm -f empirical_normal.*
	make reset
	bbk_sheet -mode mk > Makefile