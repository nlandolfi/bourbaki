# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make scalar_multiple_sets.pdf
	make graph.pdf

scalar_multiple_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > scalar_multiple_sets.tex

scalar_multiple_sets.pdf: ../../*.tex ../../trademark.pdf *.tex scalar_multiple_sets.tex
	# pdflatex --file-line-error -interaction=nonstopmode scalar_multiple_sets.tex
	../../latexrun scalar_multiple_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: scalar_multiple_sets.pdf
	open scalar_multiple_sets.pdf

o:
	make open

terms: scalar_multiple_sets.tex
	bbk_sheet -mode ts scalar_multiple_sets.tex

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
	rm -f scalar_multiple_sets.*
	make reset
	bbk_sheet -mode mk > Makefile