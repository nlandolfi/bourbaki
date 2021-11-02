# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make multiway_classification_models.pdf

multiway_classification_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multiway_classification_models.tex

multiway_classification_models.pdf: ../../*.tex ../../trademark.pdf *.tex multiway_classification_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode multiway_classification_models.tex
	../../latexrun multiway_classification_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multiway_classification_models.pdf
	open multiway_classification_models.pdf

o:
	make open

terms: multiway_classification_models.tex
	bbk_sheet -mode ts multiway_classification_models.tex

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
	rm -f multiway_classification_models.*
	make reset
	bbk_sheet -mode mk > Makefile