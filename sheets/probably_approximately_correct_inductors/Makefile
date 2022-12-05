# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probably_approximately_correct_inductors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

probably_approximately_correct_inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probably_approximately_correct_inductors.tex

probably_approximately_correct_inductors.pdf: ../../*.tex ../../trademark.pdf *.tex probably_approximately_correct_inductors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probably_approximately_correct_inductors.tex
	../../latexrun probably_approximately_correct_inductors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probably_approximately_correct_inductors.pdf
	open probably_approximately_correct_inductors.pdf

o:
	make open

terms: probably_approximately_correct_inductors.tex
	bbk_sheet -mode ts probably_approximately_correct_inductors.tex

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
	rm -f probably_approximately_correct_inductors.*
	make reset
	bbk_sheet -mode mk > Makefile