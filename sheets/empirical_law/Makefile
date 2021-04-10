# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make empirical_law.pdf
	make graph.pdf

empirical_law.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empirical_law.tex

empirical_law.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_law.tex
	# pdflatex --file-line-error -interaction=nonstopmode empirical_law.tex
	../../latexrun empirical_law.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_law.pdf
	open empirical_law.pdf

o:
	make open

terms: empirical_law.tex
	bbk_sheet -mode ts empirical_law.tex

spell:
	aspell -c empirical_law.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f empirical_law.*
	make reset
	bbk_sheet -mode mk > Makefile