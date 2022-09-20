# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make contingency_tables.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

contingency_tables.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > contingency_tables.tex

contingency_tables.pdf: ../../*.tex ../../trademark.pdf *.tex contingency_tables.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode contingency_tables.tex
	../../latexrun contingency_tables.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: contingency_tables.pdf
	open contingency_tables.pdf

o:
	make open

terms: contingency_tables.tex
	bbk_sheet -mode ts contingency_tables.tex

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
	rm -f contingency_tables.*
	make reset
	bbk_sheet -mode mk > Makefile