# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make accounts.pdf

accounts.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > accounts.tex

accounts.pdf: ../../*.tex ../../trademark.pdf *.tex accounts.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode accounts.tex
	../../latexrun accounts.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: accounts.pdf
	open accounts.pdf

o:
	make open

terms: accounts.tex
	bbk_sheet -mode ts accounts.tex

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
	rm -f accounts.*
	make reset
	bbk_sheet -mode mk > Makefile