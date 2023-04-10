# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make correlation_and_independence.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

correlation_and_independence.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > correlation_and_independence.tex

correlation_and_independence.pdf: ../../*.tex ../../trademark.pdf *.tex correlation_and_independence.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode correlation_and_independence.tex
	../../latexrun correlation_and_independence.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: correlation_and_independence.pdf
	open correlation_and_independence.pdf

o:
	make open

terms: correlation_and_independence.tex
	bbk_sheet -mode ts correlation_and_independence.tex

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
	rm -f correlation_and_independence.*
	make reset
	bbk_sheet -mode mk > Makefile