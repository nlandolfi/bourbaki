# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make statistical_models.pdf

statistical_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > statistical_models.tex

statistical_models.pdf: ../../*.tex ../../trademark.pdf *.tex statistical_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode statistical_models.tex
	../../latexrun statistical_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: statistical_models.pdf
	open statistical_models.pdf

o:
	make open

terms: statistical_models.tex
	bbk_sheet -mode ts statistical_models.tex

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
	rm -f statistical_models.*
	make reset
	bbk_sheet -mode mk > Makefile