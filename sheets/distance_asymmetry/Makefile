# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make distance_asymmetry.pdf
	make graph.pdf

distance_asymmetry.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distance_asymmetry.tex

distance_asymmetry.pdf: ../../*.tex ../../trademark.pdf *.tex distance_asymmetry.tex
	# pdflatex --file-line-error -interaction=nonstopmode distance_asymmetry.tex
	../../latexrun distance_asymmetry.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distance_asymmetry.pdf
	open distance_asymmetry.pdf

o:
	make open

terms: distance_asymmetry.tex
	bbk_sheet -mode ts distance_asymmetry.tex

spell:
	aspell -c distance_asymmetry.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f distance_asymmetry.*
	make reset
	bbk_sheet -mode mk > Makefile