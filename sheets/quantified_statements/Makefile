# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quantified_statements.pdf

quantified_statements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > quantified_statements.tex

quantified_statements.pdf: ../../*.tex ../../trademark.pdf *.tex quantified_statements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quantified_statements.tex
	../../latexrun quantified_statements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quantified_statements.pdf
	open quantified_statements.pdf

o:
	make open

terms: quantified_statements.tex
	bbk_sheet -mode ts quantified_statements.tex

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
	rm -f quantified_statements.*
	make reset
	bbk_sheet -mode mk > Makefile