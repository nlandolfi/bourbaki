# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make nonnegative_operators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

nonnegative_operators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > nonnegative_operators.tex

nonnegative_operators.pdf: ../../*.tex ../../trademark.pdf *.tex nonnegative_operators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode nonnegative_operators.tex
	../../latexrun nonnegative_operators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nonnegative_operators.pdf
	open nonnegative_operators.pdf

o:
	make open

terms: nonnegative_operators.tex
	bbk_sheet -mode ts nonnegative_operators.tex

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
	rm -f nonnegative_operators.*
	make reset
	bbk_sheet -mode mk > Makefile