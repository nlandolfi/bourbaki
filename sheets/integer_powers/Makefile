# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_powers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

integer_powers.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > integer_powers.tex

integer_powers.pdf: ../../*.tex ../../trademark.pdf *.tex integer_powers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_powers.tex
	../../latexrun integer_powers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_powers.pdf
	open integer_powers.pdf

o:
	make open

terms: integer_powers.tex
	bbk_sheet -mode ts integer_powers.tex

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
	rm -f integer_powers.*
	make reset
	bbk_sheet -mode mk > Makefile