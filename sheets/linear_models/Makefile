# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_models.pdf

linear_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_models.tex

linear_models.pdf: ../../*.tex ../../trademark.pdf *.tex linear_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_models.tex
	../../latexrun linear_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_models.pdf
	open linear_models.pdf

o:
	make open

terms: linear_models.tex
	bbk_sheet -mode ts linear_models.tex

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
	rm -f linear_models.*
	make reset
	bbk_sheet -mode mk > Makefile