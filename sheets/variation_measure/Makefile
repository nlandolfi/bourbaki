# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make variation_measure.pdf

variation_measure.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > variation_measure.tex

variation_measure.pdf: ../../*.tex ../../trademark.pdf *.tex variation_measure.tex
	# pdflatex --file-line-error -interaction=nonstopmode variation_measure.tex
	../../latexrun variation_measure.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: variation_measure.pdf
	open variation_measure.pdf

o:
	make open

terms: variation_measure.tex
	bbk_sheet -mode ts variation_measure.tex

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
	rm -f variation_measure.*
	make reset
	bbk_sheet -mode mk > Makefile