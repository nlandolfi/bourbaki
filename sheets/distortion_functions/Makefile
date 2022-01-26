# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distortion_functions.pdf

distortion_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distortion_functions.tex

distortion_functions.pdf: ../../*.tex ../../trademark.pdf *.tex distortion_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distortion_functions.tex
	../../latexrun distortion_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distortion_functions.pdf
	open distortion_functions.pdf

o:
	make open

terms: distortion_functions.tex
	bbk_sheet -mode ts distortion_functions.tex

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
	rm -f distortion_functions.*
	make reset
	bbk_sheet -mode mk > Makefile