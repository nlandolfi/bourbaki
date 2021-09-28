# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probabilistic_linear_model.pdf

probabilistic_linear_model.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probabilistic_linear_model.tex

probabilistic_linear_model.pdf: ../../*.tex ../../trademark.pdf *.tex probabilistic_linear_model.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probabilistic_linear_model.tex
	../../latexrun probabilistic_linear_model.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probabilistic_linear_model.pdf
	open probabilistic_linear_model.pdf

o:
	make open

terms: probabilistic_linear_model.tex
	bbk_sheet -mode ts probabilistic_linear_model.tex

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
	rm -f probabilistic_linear_model.*
	make reset
	bbk_sheet -mode mk > Makefile