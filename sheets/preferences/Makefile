# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make preferences.pdf

preferences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > preferences.tex

preferences.pdf: ../../*.tex ../../trademark.pdf *.tex preferences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode preferences.tex
	../../latexrun preferences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: preferences.pdf
	open preferences.pdf

o:
	make open

terms: preferences.tex
	bbk_sheet -mode ts preferences.tex

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
	rm -f preferences.*
	make reset
	bbk_sheet -mode mk > Makefile