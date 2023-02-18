# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make deductions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

deductions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > deductions.tex

deductions.pdf: ../../*.tex ../../trademark.pdf *.tex deductions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode deductions.tex
	../../latexrun deductions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: deductions.pdf
	open deductions.pdf

o:
	make open

terms: deductions.tex
	bbk_sheet -mode ts deductions.tex

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
	rm -f deductions.*
	make reset
	bbk_sheet -mode mk > Makefile